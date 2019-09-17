package lin.louis.poc.pact.petstore;

import java.util.Date;
import java.util.LinkedHashMap;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.ResponseStatus;


@ControllerAdvice
public class ErrorHandler {

	@ExceptionHandler(NullPointerException.class)
	@ResponseBody
	@ResponseStatus(HttpStatus.NOT_FOUND)
	public Map<String, Object> petNotFoundException(NullPointerException e, HttpServletRequest request) {
		Map<String, Object> errorAttributes = new LinkedHashMap<>();
		errorAttributes.put("timestamp", new Date());
		errorAttributes.put("status", HttpStatus.NOT_FOUND.value());
		errorAttributes.put("message", e.getMessage());
		errorAttributes.put("error", HttpStatus.NOT_FOUND.getReasonPhrase());
		errorAttributes.put("path", request.getRequestURI());
		return errorAttributes;
	}
}
