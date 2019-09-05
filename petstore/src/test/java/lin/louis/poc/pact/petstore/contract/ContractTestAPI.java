package lin.louis.poc.pact.petstore.contract;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.TestTemplate;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.web.server.LocalServerPort;
import org.springframework.test.context.junit.jupiter.SpringExtension;

import au.com.dius.pact.provider.junit.Provider;
import au.com.dius.pact.provider.junit.State;
import au.com.dius.pact.provider.junit.loader.PactBroker;
import au.com.dius.pact.provider.junit5.HttpTestTarget;
import au.com.dius.pact.provider.junit5.PactVerificationContext;
import au.com.dius.pact.provider.junit5.PactVerificationInvocationContextProvider;
import lin.louis.poc.pact.petstore.cat.Cat;
import lin.louis.poc.pact.petstore.cat.CatRepository;
import lin.louis.poc.pact.petstore.dog.Dog;
import lin.louis.poc.pact.petstore.dog.DogRepository;


@ExtendWith(SpringExtension.class)
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
// provider name must be unique and same as the one defined in the pact broker
@Provider("petstore")
// properties are set from Java system properties
@PactBroker
class ContractTestAPI {

	@LocalServerPort
	private int serverPort;

	@Autowired
	private CatRepository catRepository;

	@Autowired
	private DogRepository dogRepository;

	@TestTemplate
	@ExtendWith(PactVerificationInvocationContextProvider.class)
	void testTemplate(PactVerificationContext context) {
		context.verifyInteraction();
	}

	@BeforeEach
	void setUp(PactVerificationContext context) {
		context.setTarget(new HttpTestTarget("localhost", serverPort));
	}

	// CAT -------------------------------------------------------------------------------------------------------------

	@State("there is a cat with an id 88")
	void toExistingCat() {
		catRepository.deleteAll();
		catRepository.save(new Cat(88, "Grumpy cat", "Tardar Sauce"));
	}

	@State("there is no cat with an id 888")
	void toNonExistingCat() {
		catRepository.deleteAll();
	}

	@State("creating a Tardar Sauce cat whose name is Grumpy cat")
	void toAddCat() {}

	// DOG -------------------------------------------------------------------------------------------------------------

	@State("there is a dog with an id 88")
	void toExistingDog() {
		dogRepository.deleteAll();
		dogRepository.save(new Dog(88, "Chico", "Shiba Inu"));
	}

	@State("there is no dog with an id 888")
	void toNonExistingDog() {
		dogRepository.deleteAll();
	}

	@State("creating a Shiba Inu dog whose name is Chico")
	void toAddDog() {}
}
