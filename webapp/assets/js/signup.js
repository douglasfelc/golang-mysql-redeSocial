$("#form-signup").on("submit", userSignUp)

function userSignUp(event){
  // Prevent form submission
  event.preventDefault()

  if( $("#password").val() != $("#confirm-password").val() ){
    Swal.fire("Ops...", "Passwords don't match", "error")
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
        Swal.fire("Ops...", "Error registering user", "error")
      }else{
        // StatusCode: range of 200
        Swal.fire("Success!", "User registered successfully!", "success")
        .then(function() {
          $.ajax({
            url: "/signin",
            method: "POST",
            data: {
              email: $('#email').val(),
              password: $('#password').val()
            }
          }).done(function() {
            window.location = "/feed";
          }).fail(function() {
            Swal.fire("Ops...", "Error authenticating user!", "error");
          })
        })
      }
    }
  })

}
