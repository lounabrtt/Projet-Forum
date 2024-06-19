package actions

import (
	"database/sql"
	"ff/database/models"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

func InitDefaultNews(db *sql.DB) error {
	news := []models.New{
		{
			UUID:    uuid.Must(uuid.NewV4()).String(),
			Title:   "Hit me hard and soft",
			Content: "Billie Eilish, the global pop sensation, has once again captivated her fans with her latest album titled 'Hit Me Hard and Soft'. \n\n This bold and emotional opus confirms the undeniable talent of the young artist. \n\n A blend of emotions and depth 'Hit Me Hard and Soft' immerses the listener in an introspective journey through the ups and downs of life. \n\n The poignant lyrics and Billie Eilish's enchanting voice take the listener on a whirlwind of emotions, oscillating between strength and vulnerability. \n\n Innovative musical production \n\n The musical production of this album is as remarkable as Billie Eilish's vocal performances. The subtle arrangements and unique sounds create an immersive atmosphere that perfectly complements the artist's profound lyrics. \n\n In conclusion, 'Hit Me Hard and Soft' is an album that resonates with the audience, offering a captivating and immersive musical experience.  Billie Eilish continues to push the boundaries of creativity and emotion through this opus that is sure to leave a lasting impression on the minds and hearts of her fans worldwide.",
			Date:    "2024-05-28",
			Author:  "Louna",
		},
		{
			UUID:    uuid.Must(uuid.NewV4()).String(),
			Title:   "Radical Optimism",
			Content: "Dua Lipa, the international pop sensation, has enchanted her fans with the release of her latest album titled 'Radical Optimism'.  This bold and energetic opus testifies to the talent and creativity of the British singer. \n\n A message of positivity and hope \n\n 'Radical Optimism' is an album that exudes joy and positivity. The songs on the album convey a message of hope and optimism, inviting the listener to see the bright side of life and embrace each moment with optimism. \n\n Catchy and modern sounds \n\n The music of 'Radical Optimism' is a blend of catchy pop sounds and modern beats that encourage dancing. Dua Lipa explores new musical horizons while staying true to her unique style, creating a dynamic and captivating album.  \n\n In conclusion, 'Radical Optimism' is an album that radiates good vibes and celebrates life and music in all their splendor. Dua Lipa continues to charm her audience with catchy tracks and positive messages, confirming her place among the most influential artists in the current music industry.",
			Date:    "2024-05-04",
			Author:  "Maxence",
		},
		{
			UUID:    uuid.Must(uuid.NewV4()).String(),
			Title:   "American Dream",
			Content: "The American rapper 21 Savage shook the music industry with the release of his acclaimed album 'American Dream'. This powerful and introspective opus immerses the listener in the raw realities of life in the United States, offering an unfiltered look at the challenges and triumphs of the contemporary American experience. \n\n Exploration of Social Themes \n\n 'American Dream' tackles the social themes that shape today's American society in a frank and direct manner. 21 Savage delivers impactful lyrics that address street violence, social injustice, the struggle for survival, and the pursuit of success despite obstacles. \n\n Artistic Collaboration \n\n The album highlights 21 Savage's collaboration with other renowned artists in the music industry, adding additional dimensions to his compelling narrative. Featuring artists from diverse backgrounds enriches the album's sonic diversity and enhances its emotional impact. \n\n Critical and Public Reception \n\n 'American Dream' has been praised by critics for its sincerity, urban poetry, and ability to capture the essence of modern American experience. The album has also achieved commercial success, attracting a wide audience and solidifying 21 Savage's position as a prominent artist of his generation. \n\n In conclusion, 'American Dream' by 21 Savage is more than just a rap album; it is a powerful testimony to life, struggles, and dreams in the United States. Through his music, 21 Savage provides an authentic voice to the often overlooked realities of American society, making this album a must-listen for fans of conscious and socially engaged rap.",
			Date:    "2024-01-12",
			Author:  "Melvin",
		},
		{
			UUID:    uuid.Must(uuid.NewV4()).String(),
			Title:   "Blue lips",
			Content: "Schoolboy Q, the talented American rapper, recently released his sixth album titled 'Blue Lips'. This album stands out for its authenticity and eclecticism, offering listeners a unique and captivating musical experience. \n\n A raw and daring album \n\n 'Blue Lips' aims to be an even more raw and daring album than Schoolboy Q's previous projects. With sharp lyrics and innovative sounds, the artist pushes the boundaries of rap and presents a musical universe rich in emotions. \n\n Critical reception Critics praise \n\n Schoolboy Q's boldness and commitment in this album. The deep lyrics and refined production of 'Blue Lips' have managed to captivate rap and urban music fans, once again confirming the artist's talent for storytelling through his music.  \n\n In conclusion, 'Blue Lips' is an album that showcases Schoolboy Q's talent and creativity. With this album, the artist continues to make a mark in the music industry with his unique style and ability to innovate, offering listeners an unforgettable musical experience.",
			Date:    "2024-03-01",
			Author:  "Louna",
		},
		{
			UUID:    uuid.Must(uuid.NewV4()).String(),
			Title:   "Cowboy Carter",
			Content: "Beyonce, the pop music icon, once again dazzled the world with the release of her latest album titled 'Cowboy Carter'. This bold and innovative opus testifies to the singer's constant artistic evolution. \n\n An exploration of new musical horizons  \n\n 'Cowboy Carter' marks a turning point in Beyoncé's career by exploring country and folk sounds that the pop diva had not yet touched. The tracks on the album skillfully blend pop and country influences, creating a unique and captivating musical universe. \n\n Deep and engaging lyrics  \n\n The lyrics of 'Cowboy Carter' are both poetic and engaging, addressing universal themes such as love, freedom, and emancipation. Beyoncé delivers emotional vocal performances that transport the listener to the heart of the stories told in each song. \n\n In conclusion, 'Cowboy Carter' is an artistic work that testifies to Beyoncé's creative genius and her ability to constantly reinvent herself. This album offers a musical experience rich in emotions and authenticity, once again confirming Beyoncé's place among the greatest artists of her generation.",
			Date:    "2024-03-29",
			Author:  "Jeanne",
		},
	}

	for _, new := range news {
		exists, err := newExists(db, new.Title)
		if err != nil {
			return fmt.Errorf("failed to check if new exists: %v", err)
		}

		if exists {
			log.Printf("new with title %s already exists. Skipping insertion.", new.Title)
			continue
		}
		if err := CreateDefaultNews(db, new.UUID, new.Title, new.Content, new.Date, new.Author); err != nil {
			return fmt.Errorf("error inserting new %s: %v", new, err)
		}
	}

	return nil
}

func CreateDefaultNews(db *sql.DB, uuid, title, content, date, author string) error {
	_, err := db.Exec(`INSERT INTO news (UUID, title, content, date, author) VALUES (?, ?, ?, ?, ?)`, uuid, title, content, date, author)
	if err != nil {
		log.Printf("Error inserting new %s: %v", title, err)
		return fmt.Errorf("error inserting default new: %v", err)
	}
	log.Printf("Successfully inserted new %s", title)
	return nil
}

func newExists(db *sql.DB, title string) (bool, error) {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM news WHERE title = ?`, title).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error checking if user exists: %v", err)
	}
	return count > 0, nil
}