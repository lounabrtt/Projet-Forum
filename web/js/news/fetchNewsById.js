function getNewsIdFromUrl() {
  const params = new URL(document.location.toString()).searchParams;
  const id = params.get("id");

  return id;
}

function fetchNewsById() {
  const id = getNewsIdFromUrl();

  //   TODO: Replace this fake data by the real data from the API.
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

  return FAKE_ARTICLES.find((article) => article.id == id);
}

function addNewsOnPage() {
  const article = fetchNewsById();
  const $header = document.querySelector("#header-news");
  const $content = document.querySelector("#content-news");

  $header.innerHTML = `
    <h1 class="text-3xl font-bold">${article.title}</h1>
    <p class="text-gray-500 font-light">${article.date}</p>
    <p class="px-4 py-2 bg-blue-600/20 text-blue-500 w-fit rounded-full font-light">${article.category}</p>
  `;

  $content.innerHTML = article.description;
}

fetchNewsById();
addNewsOnPage();
