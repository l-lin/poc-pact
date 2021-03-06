package lin.louis.poc.pact.petstore.contract.pact;

import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import au.com.dius.pact.provider.junit.Consumer;
import au.com.dius.pact.provider.junit.Provider;
import au.com.dius.pact.provider.junit.State;
import au.com.dius.pact.provider.junit.loader.PactBroker;
import au.com.dius.pact.provider.junit.target.Target;
import au.com.dius.pact.provider.junit.target.TestTarget;
import au.com.dius.pact.provider.spring.SpringRestPactRunner;
import au.com.dius.pact.provider.spring.target.SpringBootHttpTarget;
import lin.louis.poc.pact.petstore.dog.Dog;
import lin.louis.poc.pact.petstore.dog.DogRepository;


/**
 * Testing pact on provider side to publish results.
 */
@RunWith(SpringRestPactRunner.class)
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
// provider name must be unique and same as the one defined in the pact broker
@Provider("petstore")
// consumer name must be unique and same as the one defined in the pact broker
@Consumer("doggy-cli")
// properties are set from Java system properties
@PactBroker
public class DogTestAPI {

	@TestTarget
	public final Target target = new SpringBootHttpTarget();

	@Autowired
	private DogRepository dogRepository;

	@State("there is a dog with an id 88")
	public void toExistingDog() {
		dogRepository.deleteAll();
		dogRepository.save(new Dog(88, "Chico", "Shiba Inu"));
	}

	@State("there is no dog with an id 888")
	public void toNonExistingDog() {
		dogRepository.deleteAll();
	}

	@State("creating a Shiba Inu dog whose name is Chico")
	public void toAddDog() {}
}
