$.get(window.location.origin+`/assets/data/posts/index.json`, function(response){
  const postList = $('#posts')
  $.each(response, function(id, title){
    let time = Number(id);
    var date = new Date(time * 1000);
    postList.prepend(`<li>${date.toLocaleString()} - <a href="read.html?postID=${id}">${title}</a></li>`)
  })
})