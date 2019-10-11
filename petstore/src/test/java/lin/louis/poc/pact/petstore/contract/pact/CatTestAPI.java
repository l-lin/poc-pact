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
import lin.louis.poc.pact.petstore.cat.Cat;
import lin.louis.poc.pact.petstore.cat.CatRepository;


/**
 * Testing pact on provider side to publish results.
 */
@RunWith(SpringRestPactRunner.class)
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
// provider name must be unique and same as the one defined in the pact broker
@Provider("petstore")
// consumer name must be unique and same as the one defined in the pact broker
@Consumer("kitty-cli")
// properties are set from Java system properties
@PactBroker
public class CatTestAPI {

	@TestTarget
	public final Target target = new SpringBootHttpTarget();

	@Autowired
	private CatRepository catRepository;

	@State("there is a cat with an id 88")
	public void toExistingCat() {
		catRepository.deleteAll();
		catRepository.save(new Cat(88, "Grumpy cat", "Tardar Sauce"));
	}

	@State("there is no cat with an id 888")
	public void toNonExistingCat() {
		catRepository.deleteAll();
	}

	@State("creating a Tardar Sauce cat whose name is Grumpy cat")
	public void toAddCat() {}
}
