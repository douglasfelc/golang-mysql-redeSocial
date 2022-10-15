$("#form-signin").on("submit", userSignIn)

function userSignIn(event){
  // Prevent form submission
  event.preventDefault()

  $.ajax({
    url: "/signin",
    method: "POST",
    data: {
      email: $("#email").val(),
      password: $("#password").val(),
    }
  }).done(function(){
    // StatusCode: range of 200
    console.log("Signin successfully!")
    window.location = "/feed"
  }).fail(function(){
    // StatusCode: range of 400 or 500
    Swal.fire("Error signin!", "Incorrect username or password", "error")
  })

}
