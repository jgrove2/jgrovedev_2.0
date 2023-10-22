# jgrovedev_2.0
# 2.0 Version of Personal Site
## Source Code
### GitHub:
Source code stored on GitHub: [Here](https://github.com/jgrove2/jgrovedev_2.0/tree/main)
## Why Go
Why Go? Well it came down to this I spend the better part of my week writing code in React for work. I wanted to try something new something that was not the typical Javascript heavy website design. Following in the footsteps of a few youtubers that I follow I ended up choosing to write my website solely using go templating as it would give me a good challenge.
## Styling?
I chose bootstrap 5 to do my tempting because I didn't want to have to worry about styling everything. I saw that bootstrap gave me a good way to quickly and easily implement styling into the go templates.  The main focus of this project was go learn how to write a website using templating and not to learn how to style everything in css so I didn't want to have to spend most of my time learning new css tecnhiques. 
## Hosting
I decided to host the website on a digital ocean droplet. This cost me $5 a month to host and gives me the ability to automatically deploy updates anytime the main branch is updated in GitHub. The process of uploading to the droplet was really easy and doesn't require me to do any crazy configuration of docker containers or firewalls to get it up and running.
## Storage of Data
### S3
To store all of the images on this site I just threw them into a public s3 bucket. In the future I am going to be looking into converting this over to [Backblaze](https://www.backblaze.com/cloud-storage) and do some more compression on the images to reduce the amount of storage that I use. Until then though I am just going to pray that I stay in the free tier and that I don't lose all my savings to AWS.
### PostgresSql
I wanted to use this time to work on my SQL knowledge base so I chose PostgresSQL. Searching online I fell upon Cockroachlabs a service that had a free tier of a Postgres server that was perfect for this project. I setup the differnt tables and landed on this configuration for the database schema. Its my first time ever creating a database schema so sorry if it is a little rough.
![database schema](https://s3.amazonaws.com/jgrovedev2.0bucket/1/dbrelationship.png)
## Conclusion
In the end of the day I am really only here to learn. All of the stuff I said above about my website are all subject to change. I have lots of ideas and very little experience so I am going to keep trying to learn and improve. If you read this far thanks for reading my first post on this website, I hope there will be many more to come stay tuned. 
> With sufficient thrust, pigs fly just fine. However, this is not necessarily a good idea. It is hard to be sure where they are going to land, and it could be dangerous sitting under them as they fly overhead.
