$('#unfollow').on('click', unFollow);
$('#follow').on('click', unFollow);

function unFollow() {
  if (this.id == "unfollow") {
    var action = "unfollow"
  } else {
    var action = "follow"
  }

  const userId = $(this).data('user-id');
  $(this).prop('disabled', true);

  $.ajax({
    url: "/users/"+userId+"/"+action,
    method: "POST"
  }).done(function() {
    window.location = "/users/"+userId;
  }).fail(function() {
    Swal.fire("Ops...", "Error "+action+" user!", "error");
    $("#"+action).prop('disabled', false);
  });

}