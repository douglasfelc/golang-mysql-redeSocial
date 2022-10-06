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
    },
    complete: function (response) {
      console.info(response)
      console.info(response.status)

      if (response.status >= 400){
        // StatusCode: range of 400 or 500
        alert("Error signin! Incorrect username or password!")
      }else{
        // StatusCode: range of 200
        alert("Signin successfully!")
      }
    }
  }).done(function(){
    console.log("Done signin")
    //window.location = "/"

  }).fail(function(){
    console.log("Fail signin")
  })

}
