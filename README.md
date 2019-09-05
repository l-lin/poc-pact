# Pact Foundation POC

> Consumer contract testing using [pact][]

POC testing:

- [pact][] consumer contract testing workflow
- [redoc][] tool to generate OpenAPI documentation
- [swagger-request-validator-pact][] to validate OAS against pact files

## Prerequisites

- java 11
- maven
- golang
- docker + docker-compose
- [pact standalone executables](https://github.com/pact-foundation/pact-ruby-standalone/releases)

## Participants

- [petstore][]: the API provider which is a simple SpringBoot webapp that
  offers CRUD operations on `/cats` and `/dogs` endpoints
- [kitty-cli][]: a consumer in form of a cli that attacks `/cats` APIs
- [doggy-cli][]: a consumer in form of a cli that attacks `/dogs` APIs

## Usage

Some usage examples:

- https://docs.pact.io/pact_broker#how-would-i-use-the-pact-broker
- https://github.com/pact-foundation/pact-ruby/wiki/Development-workflow

```bash
# Launch pact-broker.
docker-compose up

# Check pact broker UI.
firefox http://localhost:9292&

# ########################## kitty-cli consumer part ##########################
# The consumer kitty-cli has a new feature and wants the petstore provider
# to implement the corresponding API /cats.
# #############################################################################
pushd kitty-cli
# Build the client.
make compile
# Publish the kitty-cli consumer pact file with label "dev".
# You can set your CI to perform this action when the project is pushed to the
# remote git server.
make publish-pact-dev
# You can check if you can deploy the latest version (it should respond "no").
# Same here, you can set your CI to perform this check and give a feedback to
# the developers (by email, webhook, ...).
pact-broker can-i-deploy -a kitty-cli -b localhost:9292 --latest
# This should return 1 to indicate there is an error.
# It's an useful tool as it can be used in a CI when trying to deploy the app.
echo $1
# You can also check the result from pact-broker web interface.
firefox http://localhost:9292/pacts/provider/petstore/consumer/kitty-cli/latest&
popd

# ########################## petstore provider part ##########################
# The provider team has been informed of the need of the new API /cats
# either by the team or because the client build has triggered a webhook that
# build and test the provider against the pact file.
# They implement the API and publish petstore provider pact files.
# #############################################################################
pushd petstore
# The provider team writes the OpenAPI spec.
cat src/main/resources/static/openapi.yaml
# Build the provider project and run the provider app.
# It also perform the following under the hood:
# - swagger-cli validate src/main/resources/static/openapi.yaml
#   - tools to validate the OAS file
#   - this tool can be used by a CI in a merge request to check the spec syntax
# - redoc-cli generate src/main/resources/static/openapi.yaml --output build/static/index.html
#   - generate the OpenAPI documetation and add it to the app's classpath
#   - the OAS was included in the project for simplicity purpose, but it can also
#     be another project
mvn clean spring-boot:run > /dev/null 2>&1&
# Check the OpenAPI documentation.
firefox http://localhost:8080&
# Check the OAS file against the consumer pact files.
# /!\ The swagger-request-validator-pact does not publish the result to the pact-broker!
# It also executes the contract tests and verify the pact.
# It publishes the result back to the pact broker.
# Unfortunately, the pact-jvm-provider-junit does not support publishing tags.
mvn integration-test -DskipITs=false
# Now, when checking if we can deploy the consumer kitty-cli, the result is positive.
pact-broker can-i-deploy -a kitty-cli -b localhost:9292 --latest
# It should return 0
echo $?
popd

# ########################## doggy-cli consumer part ##########################
# The consumer doggy-cli has a new feature and wants the petstore provider
# to implement the corresponding API /dogs.
# #############################################################################
pushd doggy-cli
# Build the client.
make compile
# Publish the doggy-cli consumer pact file with label "dev".
make publish-pact-dev
# Checking that we can't deploy the new cli as the provider still did not verify
# the pact files.
pact-broker can-i-deploy -a doggy-cli -b localhost:9292 --latest
# This should return 1 to indicate there is an error.
echo $1
# You can also check the result from pact-broker web interface.
firefox http://localhost:9292/pacts/provider/petstore/consumer/doggy-cli/latest&
# You can also check the network graph of the consumers that attacks the petstore
# provider.
firefox http://localhost:9292/groups/petstore&
popd

# ########################## kitty-cli consumer part ##########################
# The consumer kitty-cli has developed all its features. So now, it needs to
# publish its pact files with the prod label.
# #############################################################################
pushd kitty-cli
make publish-pact-to-prod
# You can see it automatically test with the provider pact results.
firefox http://localhost:9292/matrix/provider/petstore/consumer/kitty-cli&
# You can see kitty-cli can still be deployed
pact-broker can-i-deploy -a kitty-cli -b localhost:9292 --latest
# It should return 0
echo $1
popd

# ########################## petstore provider part ##########################
# The provider team has been informed of the need of the new API /dogs
# either by the team or because the client build has triggered a webhook that
# build and test the provider against the pact file.
# They implement the API and publish petstore provider pact files.
# #############################################################################
pushd petstore
mvn integration-test -DskipITs=false
# Now that the provider has verified the contract given by the doggy-cli team,
# the doggy-cli can be deployed.
pact-broker can-i-deploy -a doggy-cli -b localhost:9292 --latest
# It should return 0
echo $?
popd

# ########################## cleanup ##########################
fg %1
^C
docker-compose down
```

## Notes

This POC is not using the latest version of pact-jvm (which is 4.0.0-beta6 at
the time the POC was performed) because this POC is also testing the library
[swagger-request-validator-pact][] which uses the [3.5.20][pact-version].

[pact-jvm-provider-junit][] does not support publishing provider pact results
with a tag (e.g. set a `dev` or `prod` to a provider pact result), which limits
the usage of pact. See the discussions [Not able to publish tags to PACT Broker
from Consumer JUnit Test][] and [Publish verification results with a version
tag in pact-jvm-provider-spring_2.12][].
You can however update and add the tag to the pact result
afterward according to [bethesque][update pact result] using a `PUT` to a
pact broker endpoint or using the pact-broker-client.

[pact-jvm-provider-spring][] does not support JUnit5.

[pact-jvm-provider-maven][] needs a running environment of the provider to
test the pact files. Thus, it's necessarily to either:

- mount a special environment, like for end to end tests, but one for contract
  tests
- mount only the app when performing the contract tests and destroy it after the
  tests
  - the app must mock all its external interactions

[pact]: https://docs.pact.io/
[redoc]: https://redocly.github.io/redoc/
[petstore]: petstore
[kitty-cli]: kitty-cli
[doggy-cli]: doggy-cli
[swagger-request-validator-pact]: https://bitbucket.org/atlassian/swagger-request-validator/src/master/swagger-request-validator-pact/
[pact-version]: https://bitbucket.org/atlassian/swagger-request-validator/src/d151bff4702ab00e939c9b75fd1f41c5bc0215a7/pom.xml#lines-65
[pact-jvm-provider-junit]: https://github.com/DiUS/pact-jvm/tree/3_5_20/pact-jvm-provider-junit
[pact-jvm-provider-spring]: https://github.com/DiUS/pact-jvm/tree/3_5_20/pact-jvm-provider-spring
[pact-jvm-provider-maven]: https://github.com/DiUS/pact-jvm/tree/3_5_20/pact-jvm-provider-maven
[update pact result]: https://github.com/DiUS/pact-jvm/issues/823#issuecomment-443895021
[Not able to publish tags to PACT Broker from Consumer JUnit Test]: https://github.com/DiUS/pact-jvm/issues/459
[Publish verification results with a version tag in pact-jvm-provider-spring_2.12]: https://github.com/DiUS/pact-jvm/issues/823