package lin.louis.poc.pact.petstore.cat;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

import lin.louis.poc.pact.petstore.PetNoFoundException;


@RestController
@RequestMapping(path = "/cats")
public class CatController {

	private final CatRepository catRepository;

	@Autowired
	public CatController(CatRepository catRepository) {this.catRepository = catRepository;}

	@GetMapping(path = "/{id}")
	public Cat get(@PathVariable long id) {
		return catRepository.findById(id)
				.orElseThrow(() -> new PetNoFoundException(id));
	}

	@PostMapping
	@ResponseStatus(HttpStatus.CREATED)
	public Cat save(@RequestBody Cat cat) {
		return catRepository.save(cat);
	}
}
