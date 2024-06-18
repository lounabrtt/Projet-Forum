const FAKE_POSTS = [
  {
    id: 1,
    title: "Post 1",
    content: "This is post 1",
    userId: 1,
  },
  {
    id: 2,
    title: "Post 2",
    content: "This is post 2",
    userId: 1,
  },
];

function renderAllPosts() {
  const $container = document.querySelector("#posts");

  return FAKE_POSTS.forEach((post) => {
    $container.innerHTML += `<a href="/post?id=${post.id}" class="flex flex-col gap-2 max-w-[420px] p-4 rounded-md border border-gray-600 transition-all hover:transform hover:-translate-y-2">
      <h3 class="font-medium text-xl">${post.title}</h3>
      <p class="text-gray-400">${post.content}</p>
    </a>`;
  });
}

renderAllPosts();
