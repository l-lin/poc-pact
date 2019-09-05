package lin.louis.poc.pact.petstore.contract;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.web.server.LocalServerPort;
import org.springframework.test.context.junit.jupiter.SpringExtension;

import com.atlassian.oai.validator.pact.PactProviderValidationResults;
import com.atlassian.oai.validator.pact.PactProviderValidator;


@ExtendWith(SpringExtension.class)
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
class OpenAPITestAPI {

	@LocalServerPort
	private int serverPort;

	@Test
	void validate() {
		final PactProviderValidator validator = PactProviderValidator
				.createFor("http://localhost:" + serverPort + "/openapi.yaml")
				.withPactsFrom("http://localhost:9292", "petstore")
				.build();
		PactProviderValidationResults results = validator.validate();
		if (results.hasErrors()) {
			Assertions.fail(
					"Validation errors found.\n\t" + results.getValidationFailureReport().replace("\n", "\n\t"));
		}
	}

}
