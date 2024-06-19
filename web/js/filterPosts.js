function filterPosts() {
  const buttons = document.querySelectorAll("#category-button");
  const postItems = document.querySelectorAll("#post-item");

  buttons.forEach((button) => {
    button.addEventListener("click", () => {
      const category = button.getAttribute("data-category");

      postItems.forEach((item) => {
        if (
          category === "all" ||
          item.getAttribute("data-category") === category
        ) {
          item.style.display = "block";
        } else {
          item.style.display = "none";
        }
      });

      buttons.forEach((btn) =>
        btn.classList.remove(
          "bg-blue-500",
          "text-white",
          "border",
          "border-blue-500"
        )
      );

      button.classList.add(
        "bg-blue-500",
        "text-white",
        "border",
        "border-blue-500"
      );
    });
  });
}

filterPosts();
