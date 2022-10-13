$("#unfollow").on("click", unFollow);
$("#follow").on("click", unFollow);
$("#update-profile").on("submit", updateProfile);
$("#update-password").on("submit", updatePassword);
$("#delete-user").on("click", deleteUser);

function unFollow() {
  if (this.id == "unfollow") {
    var action = "unfollow"
  } else {
    var action = "follow"
  }

  const userId = $(this).data("user-id");
  $(this).prop("disabled", true);

  $.ajax({
    url: "/users/"+userId+"/"+action,
    method: "POST"
  }).done(function() {
    window.location = "/users/"+userId;
  }).fail(function() {
    Swal.fire("Ops...", "Error "+action+" user!", "error");
    $("#"+action).prop("disabled", false);
  });

}

// updateProfile updates the information of the logged in user
function updateProfile(event) {
  // Prevent form submission
  event.preventDefault();

  $.ajax({
    url: "/update-profile",
    method: "PUT",
    data: {
      name: $("#name").val(),
      nick: $("#nick").val(),
      email: $("#email").val(),
    }
  }).done(function() {
    Swal.fire("Success!", "User successfully updated!", "success")
    .then(function() {
      window.location = "/profile";
    });
  }).fail(function() {
    Swal.fire("Ops...", "Error updating user!", "error");
  });
}

// updatePassword update the password of the logged in user
function updatePassword(event) {
  // Prevent form submission
  event.preventDefault();

  // Checks that the passwords entered match
  if ($("#new_password").val() != $("#password_confirmation").val()) {
    Swal.fire("Ops...", "Passwords don't match!", "warning");
    return;
  }

  $.ajax({
    url: "/update-password",
    method: "POST",
    data: {
      current: $("#current_password").val(),
      new: $("#new_password").val()
    }
  }).done(function() {
    Swal.fire("Success!", "Password has been updated successfully!", "success")
    .then(function() {
      window.location = "/profile";
    })
  }).fail(function() {
    Swal.fire("Ops...", "Error updating password!", "error");
  });
}

// deleteUser remove logged in user account
function deleteUser() {
  Swal.fire({
    title: "Attention!",
    text: "Are you sure you want to delete your account? This is an irreversible action!",
    showCancelButton: true,
    cancelButtonText: "Cancel",
    icon: "warning"
  }).then(function(confirmacao) {
    if (confirmacao.value) {
      $.ajax({
        url: "/delete-user",
        method: "DELETE"
      }).done(function() {
        Swal.fire("Success!", "Your user has been successfully deleted!", "success")
        .then(function() {
          window.location = "/signout";
        })
      }).fail(function() {
        Swal.fire("Ops...", "There was an error deleting the user!", "error");
      });
    }
  })
}