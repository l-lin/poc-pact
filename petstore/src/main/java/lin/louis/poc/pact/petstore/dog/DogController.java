package lin.louis.poc.pact.petstore.dog;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;


@RestController
@RequestMapping(path = "/dogs")
public class DogController {

	private final DogRepository dogRepository;

	@Autowired
	public DogController(DogRepository dogRepository) {this.dogRepository = dogRepository;}

	@GetMapping(path = "/{id}")
	public Dog get(@PathVariable long id) {
		return dogRepository.findById(id)
				.orElseThrow(() -> new NullPointerException("Dog not found for id " + id));
	}

	@PostMapping
	@ResponseStatus(HttpStatus.CREATED)
	public Dog save(@RequestBody Dog dog) {
		return dogRepository.save(dog);
	}
}
