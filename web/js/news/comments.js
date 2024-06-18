// FAKE DATA
const FAKE_COMMENTS = [
  {
    id: 1,
    author: "John Doe",
    content: "This is a fake comment",
    date: "2022-01-01",
    likes: 10,
  },
  {
    id: 2,
    author: "Jane Doe",
    content: "This is another fake comment",
    date: "2022-01-02",
    likes: 3,
  },
  {
    id: 3,
    author: "John Doe",
    content:
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit ullamcorper dictum. Lorem ipsum dolor sit amet, consectetur adipiscing elit ullamcorper dictum. Lorem ipsum dolor sit amet, consectetur adipiscing elit ullamcorper dictum.",
    date: "2022-01-03",
    likes: 300,
  },
];

function heartSvg(outlined = false) {
  if (outlined) {
    return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2c-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/></svg>`;
  }

  return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2c-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/></svg>`;
}

const $commentCount = document.querySelector("#comment-count");
const $commentsList = document.querySelector("#comments-list");

function renderComments() {
  $commentCount.textContent = `(${FAKE_COMMENTS.length})`;

  FAKE_COMMENTS.forEach((comment) => {
    const userAlreadyLiked = false;

    $commentsList.innerHTML += `<div class="flex flex-col gap-2 border-b border-gray-800 pb-6 text-sm">
    <div class="w-full justify-between flex items-center gap-4">
    <p class="font-medium">${comment.author} <span class="font-normal text-sm text-gray-500">${comment.date}</span></p>

    <button class="flex gap-2 p-2 rounded-md hover:bg-gray-800">${comment.likes} ${heartSvg(userAlreadyLiked ? "" : true)}</button>

    </div>
      <p class="text-gray-500 font-light max-w-[60ch]">${comment.content}</p>
    </div>`;
  });
}

renderComments();
