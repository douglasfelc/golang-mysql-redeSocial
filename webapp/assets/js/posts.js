$("#new-post").on("submit", newPost)

function newPost(event){
  // Prevent form submission
  event.preventDefault()

  $.ajax({
    url: "/posts",
    method: "POST",
    data: {
      content: $("#content").val(),
    },
    complete: function (response) {
      console.info(response)
      console.info(response.status)

      if (response.status >= 400){
        // StatusCode: range of 400 or 500
        alert("Error post")
      }else{
        // StatusCode: range of 200
        console.log("Post successfully!")
        window.location = "/feed"
      }
    }
  }).done(function(){
    console.log("Done post")
    //window.location = "/feed"

  }).fail(function(){
    console.log("Fail post")
  })
}