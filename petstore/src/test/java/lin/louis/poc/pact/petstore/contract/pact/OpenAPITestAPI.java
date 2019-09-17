package lin.louis.poc.pact.petstore.contract.pact;

import org.junit.Assert;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.web.server.LocalServerPort;
import org.springframework.test.context.junit4.SpringRunner;

import com.atlassian.oai.validator.pact.PactProviderValidationResults;
import com.atlassian.oai.validator.pact.PactProviderValidator;


/**
 * Test the swagger-request-validator-pact (https://bitbucket.org/atlassian/swagger-request-validator/src/master/swagger-request-validator-pact/)
 * against pact broker.
 *
 * It tests whether we can use the OAS along with pact.
 */
@RunWith(SpringRunner.class)
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
public class OpenAPITestAPI {

	@LocalServerPort
	private int serverPort;

	@Test
	public void validate() {
		final PactProviderValidator validator = PactProviderValidator
				.createFor("http://localhost:" + serverPort + "/openapi.yaml")
				.withPactsFrom("http://localhost:9292", "petstore")
				.build();
		PactProviderValidationResults results = validator.validate();
		if (results.hasErrors()) {
			Assert.fail(
					"Validation errors found.\n\t" + results.getValidationFailureReport().replace("\n", "\n\t"));
		}
	}

}
