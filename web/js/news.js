const CATEGORIES_FILTERS_TAGS = ["Pop", "Rap US", "Rap FR", "Rnb", "Soul"];
let FILTERS = [];

// ARTICLES

// FAKE DATA TO REMOVE AFTER FETCH THE REAL DATA FROM API
const FAKE_ARTICLES = [
  {
    id: 1,
    title: "Angèle sort un nouvel album 'Nonante-Cinq'",
    description:
      "La chanteuse belge Angèle a récemment sorti son nouvel album intitulé 'Nonante-Cinq'. Cet album, qui fait suite à son premier opus 'Brol', explore des thèmes variés tels que l'amour, la société et l'introspection. Angèle continue de marquer la scène musicale avec sa voix douce et ses paroles incisives, confirmant son statut d'artiste incontournable en Francophonie.",
    date: "2022-01-01",
    category: "Pop",
  },
  {
    id: 2,
    title: "Julien Doré en tournée avec 'Aimée'",
    description:
      "Après le succès de son album 'Aimée', Julien Doré a entamé une tournée nationale. Son album, dédié à sa grand-mère, aborde des sujets personnels et sociétaux avec une touche poétique et mélodique caractéristique de l'artiste. Les concerts de Julien Doré sont très attendus, promettant des performances émouvantes et intimes.",
    date: "2023-01-01",
    category: "Pop",
  },
  {
    id: 3,
    title: "Vianney collabore avec Ed Sheeran",
    description:
      "Vianney a récemment dévoilé une collaboration avec l’artiste international Ed Sheeran sur la chanson 'Call on Me'. Cette collaboration unique mélange le style folk-pop de Vianney avec la touche pop acoustique d’Ed Sheeran, créant une chanson douce et entraînante. Ce duo a été accueilli avec enthousiasme par les fans des deux artistes.",
    date: "2023-01-02",
    category: "Pop",
  },
  {
    id: 4,
    title: "Stromae revient avec 'Santé'",
    description:
      "Stromae a fait un retour très attendu sur la scène musicale avec son single 'Santé'. La chanson célèbre les travailleurs de l'ombre qui maintiennent le monde en marche, avec un rythme entraînant et des paroles pleines de gratitude. Stromae, connu pour ses textes profonds et ses mélodies accrocheuses, prouve encore une fois son talent inégalé.",
    date: "2023-01-03",
    category: "Pop",
  },
];

function addArticles() {
  const $articles = document.querySelector("#articles");
  $articles.innerHTML = "";

  const articlesWithFilter = FAKE_ARTICLES.filter((article) => {
    if (!FILTERS.length) return true;

    return FILTERS.includes(article.category);
  });

  if (!articlesWithFilter.length) {
    $articles.innerHTML = "No articles found";
    return;
  }

  articlesWithFilter.forEach((article) => {
    $articles.innerHTML += `<a href="/article?id=${article.id}" class="flex flex-col pb-4 gap-4 border-b border-gray-400 hover:transform hover:translate-x-4 transition-all duration-300">
      <h3 class="text-2xl font-medium">${article.title}</h3>
      <p class="text-gray-500 font-light">${article.description}</p>
      <p class="text-gray-500 font-light">${article.date}</p>
    </a>`;
  });
}

// CATEGORIES

function addCategoriesTags() {
  const $categories = document.querySelector("#filter-categories");

  CATEGORIES_FILTERS_TAGS.forEach((category) => {
    const $button = document.createElement("button");
    $button.textContent = category;
    $button.classList.add(
      "px-3",
      "py-1",
      "rounded-full",
      "border",
      "border-blue-500",
      "hover:bg-blue-600",
      "hover:text-white",
      "text-sm"
    );
    $button.addEventListener("click", () => {
      toggleCategoryFilter(category);
      addArticles();
    });
    $categories.appendChild($button);
  });
}

function toggleCategoryFilter(category) {
  const index = FILTERS.indexOf(category);

  if (index === -1) {
    FILTERS.push(category);
  } else {
    FILTERS.splice(index, 1);
  }

  updateButtonStyles();
}

function updateButtonStyles() {
  const $buttons = document.querySelectorAll("#filter-categories button");
  $buttons.forEach(($button) => {
    const category = $button.textContent.trim();
    if (FILTERS.includes(category)) {
      $button.classList.add("bg-blue-600", "text-white");
    } else {
      $button.classList.remove("bg-blue-600", "text-white");
    }
  });
}

//
addCategoriesTags();
addArticles();
