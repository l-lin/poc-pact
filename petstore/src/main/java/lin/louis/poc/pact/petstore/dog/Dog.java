package lin.louis.poc.pact.petstore.dog;

import javax.persistence.Entity;
import javax.persistence.Id;


@Entity
public class Dog {

	@Id
	private long id;

	private String name;

	private String type;

	public Dog() {}

	public Dog(long id, String name, String type) {
		this.id = id;
		this.name = name;
		this.type = type;
	}

	public long getId() {
		return id;
	}

	public void setId(long id) {
		this.id = id;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public String getType() {
		return type;
	}

	public void setType(String type) {
		this.type = type;
	}
}
