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
    }
  }).done(function() {
    // StatusCode: range of 200
    alert("User registered successfully!")

  }).fail(function(error) {
    console.log(error)
    // StatusCode: range of 400 or 500
    alert("Error registering user!")
  })
}
