# petstore

> Simple webapp that uses Pact as a provider

## Getting started

```bash
# Build
mvn clean package
# Execute pact verification
mvn integration-test -DskipITs=false
```

## Usage

```bash
# Run the webapp
java -jar target/*.jar
# Or using maven
mvn spring-boot:run
```
