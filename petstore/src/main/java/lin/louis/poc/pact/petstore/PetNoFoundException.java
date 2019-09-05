package lin.louis.poc.pact.petstore;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;


@ResponseStatus(HttpStatus.NOT_FOUND)
public class PetNoFoundException extends RuntimeException {
	public PetNoFoundException(long id) {
		super("No pet found with id " + id);
	}
}
