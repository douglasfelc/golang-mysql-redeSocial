$("#form-signup").on("submit", userSignUp)

function userSignUp(event){
  // Prevent form submission
  event.preventDefault()

  if( $("#password").val() != $("#confirm-password").val() ){
    alert("Passwords don't match")
    return
  }

  $.ajax({
    url: "/users",
    method: "POST",
    data: {
      name: $("#name").val(),
      nick: $("#nick").val(),
      email: $("#email").val(),
      password: $("#password").val(),
    },
    success: function (response) {
      console.log("success");
      console.info(response)
    },
    error: function (response) {
      console.log("error");
      console.info(response)
    },
    complete: function (response) {
      console.info(response)
      console.info(response.status)

      if (response.status >= 400){
        // StatusCode: range of 400 or 500
        alert("Error registering user!")
      }else{
        // StatusCode: range of 200
        alert("User registered successfully!")
      }
    }
  })

}
