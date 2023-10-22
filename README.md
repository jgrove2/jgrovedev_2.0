# 2.0 Version of Personal Site
## Source Code
### GitHub Repository:
The source code is stored on GitHub: [Here](https://github.com/jgrove2/jgrovedev_2.0/tree/main)
## Why Go
I chose Bootstrap 5 to handle the styling because I didn't want to worry about styling everything from scratch. Bootstrap offered a quick and easy way to implement styling into the Go templates. The main focus of this project was to learn how to create a website using templating, not to learn how to style everything in CSS. I didn't want to spend most of my time learning new CSS techniques.
## Styling?
I chose bootstrap 5 to do my tempting because I didn't want to have to worry about styling everything. I saw that bootstrap gave me a good way to quickly and easily implement styling into the go templates.  The main focus of this project was go learn how to write a website using templating and not to learn how to style everything in css so I didn't want to have to spend most of my time learning new css tecnhiques. 
## Posts
For the posts, I store all of the text as Markdown in a PostgreSQL table. This data is parsed using the Go package [gomarkdown](https://github.com/gomarkdown/markdown). Gomarkdown parses the Markdown into HTML, allowing it to be displayed properly in the DOM. This process also leverages the Go templating engine, as the code is parsed into a Go template and then inserted onto the screen. Here's an example of the code:
```go
func MdToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
```
The code above takes a byte string of Markdown and converts it into HTML.
## Hosting
I decided to host the website on a Digital Ocean droplet. It costs me $5 a month to host and allows me to automatically deploy updates anytime the main branch is updated on GitHub. The process of uploading to the droplet was straightforward and didn't require complex configuration of Docker containers or firewalls to get it up and running.
## Storage of Data
### S3
To store all the images on this site, I placed them in a public S3 bucket. In the future, I plan to explore converting this to [Backblaze](https://www.backblaze.com/cloud-storage) and implementing further image compression to reduce storage usage. Until then, I hope to remain in the free tier and not deplete my savings with AWS.
### PostgresSql
I wanted to use this opportunity to enhance my SQL knowledge, so I chose PostgreSQL. While searching online, I came across Cockroachlabs, a service with a free tier of a PostgreSQL server that was perfect for this project. I set up different tables and settled on this configuration for the database schema. It's my first time creating a database schema, so I apologize if it's a bit rough.
![database schema](https://s3.amazonaws.com/jgrovedev2.0bucket/1/dbrelationship.png)
## Conclusion
In the end, I'm here to learn. Everything I mentioned above about my website is subject to change. I have many ideas and limited experience, so I'll keep trying to learn and improve. If you've read this far, thank you for reading my first post on this website. I hope there will be many more to come. Stay tuned.
#### Quote
With sufficient thrust, pigs fly just fine. However, this is not necessarily a good idea. It is hard to be sure where they are going to land, and it could be dangerous sitting under them as they fly overhead.
- R. Callon [Link] (https://www.janko.at/Humor/RFC/rfc1925.htm#:~:text=(3)%20With%20sufficient%20thrust%2C,nor%20understood%20unless%20experienced%20firsthand)
