// read post id from url
let postID = $_GET("postID");
// make api call to post url and display post page
$.get(window.location.origin+`/assets/data/posts/${postID}.json`, function (response)
{
  console.log(response);
  let time = Number(postID);
  var date = new Date(time * 1000);
  // todo display loading page
  $('#post-time').text(date.toLocaleString());
  $('#post-title').text(response.title);
  $('#post-body').html(marked.parse(response.body));
  // todo hide loading page
});

function $_GET(name, url)
{
  if (!url)
  {
    url = window.location.href;
  }
  name = name.replace(/[\[\]]/g, "\\$&");
  var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
    results = regex.exec(url);
  if (!results) return null;
  if (!results[2]) return '';
  return decodeURIComponent(results[2].replace(/\+/g, " "));
}