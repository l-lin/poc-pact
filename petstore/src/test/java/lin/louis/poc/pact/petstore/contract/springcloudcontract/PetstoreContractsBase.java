package lin.louis.poc.pact.petstore.contract.springcloudcontract;

import org.junit.Before;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.web.context.WebApplicationContext;

import io.restassured.module.mockmvc.RestAssuredMockMvc;
import lin.louis.poc.pact.petstore.cat.Cat;
import lin.louis.poc.pact.petstore.cat.CatRepository;
import lin.louis.poc.pact.petstore.dog.Dog;
import lin.louis.poc.pact.petstore.dog.DogRepository;


@RunWith(SpringRunner.class)
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
public abstract class PetstoreContractsBase {

	@Autowired
	private WebApplicationContext context;

	@Autowired
	private CatRepository catRepository;

	@Autowired
	private DogRepository dogRepository;

	@Before
	public void setUp() {
		catRepository.deleteAll();
		catRepository.save(new Cat(88, "Grumpy cat", "Tardar Sauce"));
		dogRepository.deleteAll();
		dogRepository.save(new Dog(88, "Chico", "Shiba Inu"));
		RestAssuredMockMvc.webAppContextSetup(context);
	}
}
