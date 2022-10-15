$("#new-post").on("submit", newPost)
$("#update-post").on("submit", updatePost)

$(document).on("click", ".like-post", disLikePost);
$(document).on("click", ".dislike-post", disLikePost);
$(document).on("click", ".delete-post", deletePost);

function newPost(event){
  // Prevent form submission
  event.preventDefault()

  $.ajax({
    url: "/posts",
    method: "POST",
    data: {
      content: $("#content").val(),
    }
  }).done(function(){
    console.log("Post successfully!")
    window.location = "/feed"
  }).fail(function(){
    // StatusCode: range of 400 or 500
    Swal.fire("Ops...", "Error post", "error")
  })
}

function disLikePost(event) {
  // Prevent form submission
  event.preventDefault()

  // Checks if the element has the dislike class
  var hasClass = Array.from(event.target.classList).indexOf("dislike-post") > -1;
  if (hasClass == true) {
    var action = "dislike"
  } else {
    var action = "like"
  }

  // Get the clicked element
  let clickedElement = $(event.target)

  // Search for the closest div, and get the content of data-post-id
  let postId = clickedElement.closest("div").data("post-id")

  clickedElement.prop("disabled", true)
  $.ajax({
    url: "/posts/"+postId+"/"+action,
    method: "POST",
    data: {
      content: $("#content").val(),
    }
  }).done(function(){

    // Get the next span element it finds, where the current amount of likes are
    let likesCount_span = clickedElement.next("span")

    // Convert span text to integer
    let likesCount = parseInt(likesCount_span.text())

    if( action == 'like' ){
      // Change span text with incremented counter
      likesCount_span.text(likesCount + 1)

      clickedElement.addClass("text-danger")

      // Change class to on click again dislike post
      clickedElement.removeClass("like-post")
      clickedElement.addClass("dislike-post")
    }else{
      // Change span text with decremented counter
      likesCount_span.text(likesCount - 1)

      clickedElement.removeClass("text-danger")

      // Change class to on click again like post
      clickedElement.removeClass("dislike-post")
      clickedElement.addClass("like-post")
    }

  }).fail(function(response){
    Swal.fire("Ops...", "Error "+action+" post", "error")
    console.info(response)

  }).always(function() {
    clickedElement.prop("disabled", false)
  })
}

function updatePost(event){
  // Prevent form submission
  event.preventDefault()

  $("#update-post-submit-btn").prop("disabled", true)
  const postId = $('#id').val()

  $.ajax({
    url: "/posts/"+postId,
    method: "PUT",
    data: {
      content: $('#content').val()
    }
  }).done(function() {
    Swal.fire('Success!', 'Post successfully updated!', 'success')
    .then(function() {
      window.location = "/feed";
    })
  }).fail(function() {
      Swal.fire("Ops...", "Error updating post!", "error");
  }).always(function() {
    $("#update-post-submit-btn").prop("disabled", false)
  })

}

function deletePost(event) {
  event.preventDefault();

  Swal.fire({
    title: "Attention!",
    text: "Are you sure you want to delete this post? This action is irreversible!",
    showCancelButton: true,
    cancelButtonText: "Cancel",
    icon: "warning"
  }).then(function(confirm) {
    if (!confirm.value) return;

    // Get the clicked element
    const clickedElement = $(event.target)
    
    // Search for the closest div, and get the content of data-post-id
    const post = clickedElement.closest('div')
    const postId = post.data('post-id');

    clickedElement.prop('disabled', true);

    $.ajax({
      url: "/posts/"+postId,
      method: "DELETE"
    }).done(function() {
      post.fadeOut("slow", function() {
        $(this).remove();
      });
    }).fail(function() {
      Swal.fire("Ops...", "Error deleting post!", "error");
    });
  })

}