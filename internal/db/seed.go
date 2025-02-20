package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"example.com/social-app/internal/store"
)

var usernames = []string{
	"rakibul", "waell", "sabit", "anan", "salman",
	"fahim", "tariq", "junaid", "kamal", "rayan",
	"nabil", "hasib", "zaid", "adil", "omar",
	"yasin", "sajid", "karim", "tanvir", "farhan",
	"ismail", "hassan", "ayman", "rashid", "bilal",
}

var titles = []string{
	"The Future of AI", "Breaking into Tech", "Top 10 Coding Tips", "Why Go is Awesome",
	"How to Stay Productive", "The Rise of Startups", "Secrets of Successful People",
	"Mastering Data Structures", "The Power of Networking", "The Ultimate Fitness Guide",
	"Side Hustles That Work", "How to Learn Faster", "Investing for Beginners",
	"Why Remote Work is the Future", "Time Management Hacks", "The Science of Habits",
	"How to Start a Business", "Building Scalable Apps", "The Psychology of Success",
	"Freelancing vs Full-Time Jobs", "The Impact of Social Media", "How to Code Efficiently",
	"The Best Books for Entrepreneurs", "Why Minimalism Works", "The Truth About Motivation",
}

var content = []string{
	"This article explores the impact of AI on our daily lives and future careers.",
	"Breaking into the tech industry can be challenging. Here are some key tips to get started.",
	"Want to improve your coding skills? Check out these 10 essential tips.",
	"Go is one of the most efficient programming languages. Learn why it’s worth using.",
	"Struggling with productivity? These strategies will help you get more done.",
	"Startups are taking over the business world. Here’s how they’re changing industries.",
	"Success isnt an accident. Learn the habits of high achievers.",
	"Understanding data structures can level up your coding game. Here's how to master them.",
	"Networking can open doors to new opportunities. Learn how to do it effectively.",
	"Want to get in shape? This fitness guide will help you create a sustainable routine.",
	"Side hustles can be life-changing. Here are some ideas to get started.",
	"Learning efficiently is a skill. These techniques will help you learn faster.",
	"Investing can seem complex. This beginner’s guide simplifies the process.",
	"Remote work is becoming the new normal. Here's why it's here to stay.",
	"Time management can make or break your success. Master these hacks to stay ahead.",
	"Habits shape our lives. Learn how to build good habits and break bad ones.",
	"Dreaming of starting a business? Follow these steps to make it happen.",
	"Scalability is key for any app. Learn how to build software that grows with demand.",
	"Psychology plays a huge role in success. Understand the mindset of winners.",
	"Should you freelance or work full-time? This article weighs the pros and cons.",
	"Social media influences our lives more than we realize. Here’s its real impact.",
	"Writing clean, efficient code is crucial. Follow these best practices.",
	"Entrepreneurs read a lot. These books can change the way you think about business.",
	"Minimalism isn’t just about less clutter. Here’s why it leads to a better life.",
	"Motivation fades, but discipline lasts. Learn the truth about staying motivated.",
}

var comments = []string{
	"This is an interesting take! Never thought about it this way.",
	"Great article! Really helped me understand the topic better.",
	"I completely agree with this. Well said!",
	"Not sure I agree, but I appreciate the perspective.",
	"Can you provide more details on this? Sounds intriguing.",
	"This was very insightful. Thanks for sharing!",
	"I've been struggling with this topic, and this really helped!",
	"Amazing work! Keep it up.",
	"This changed my view on the subject. Thanks for posting!",
	"Does anyone have additional resources on this?",
	"Really well-written and easy to follow!",
	"Interesting, but I think there's more to consider here.",
	"I love how you explained this concept so clearly.",
	"Wow, I never knew this before! Thanks for the info.",
	"Super helpful! I’ll definitely try this out.",
	"Thanks for breaking it down in a simple way.",
	"Great perspective! More people should read this.",
	"I have a different opinion, but I respect your argument.",
	"Appreciate the effort that went into this!",
	"This is exactly what I was looking for. Thanks!",
	"Brilliant explanation! Cleared up a lot of confusion for me.",
	"I’ve been thinking about this for a while, and this article confirmed my thoughts.",
	"Can you expand on this point? I'm really curious.",
	"Nice write-up! Looking forward to more content like this.",
	"This was an eye-opener! Learned something new today.",
}

var tags = []string{
	"technology", "coding", "AI", "productivity", "startup", "business", "success", "health",
	"fitness", "learning", "investing", "remote work", "time management", "psychology", "self-improvement",
	"entrepreneurship", "software", "growth", "motivation", "habits", "networking", "side hustle",
	"career", "books", "minimalism",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		err := store.Users.Create(ctx, user)
		if err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	log.Println("Seeding completed successfully.")
}

func generateUsers(n int) []*store.User {
	users := make([]*store.User, n)

	for i := 0; i < n; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "randompass",
		}
	}

	return users
}

func generatePosts(n int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, n)

	for i := 0; i < n; i++ {
		user := users[rand.Intn(len(users))]
		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: content[rand.Intn(len(content))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(n int, users []*store.User, posts []*store.Post) []*store.Comment {
	commentStore := make([]*store.Comment, n)
	for i := 0; i < n; i++ {
		post := posts[rand.Intn(len(posts))]
		user := users[rand.Intn(len(users))]
		commentStore[i] = &store.Comment{
			Content: comments[rand.Intn(len(comments))],
			UserID:  user.ID,
			PostID:  post.ID,
		}
	}

	return commentStore
}
