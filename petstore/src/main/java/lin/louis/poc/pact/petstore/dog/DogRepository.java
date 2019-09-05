package lin.louis.poc.pact.petstore.dog;

import org.springframework.data.repository.CrudRepository;


public interface DogRepository extends CrudRepository<Dog, Long> {
}
