# Inspiration
Logan, Gabe, and I (Nate) all share an AP Calculus class - in which one of our shared friends loves to trade stocks. Due to this, we wanted to try some ourselves, but all being under eighteen - we couldn't. Paper trading, practicing with stocks without money, is unfortunatly restricted to people who use apps where they can also trade real stocks (limiting it to being over eighteen), or paid software.
Our team is all fans of enabling people to learn and better themselves, regardless of age - so this was a clear problem to us. We want an app where we can trade stocks ourselves - learning how to make money with investments - but that wouldn't be original. "Gameifying your life" is a tradition that exists in various industrys - and using inspiration from how Pokemon GO's leaderboard enabled many people to go on walks and get physically active, we decided that this app should also be a game.
# What CashCraft Does
CashCraft enables anyone to practice trading stocks - a valuable life skill in a world where the average cost of living has rapidly grown - especially compared to the average salary. It also encourages everyone to actively work on it due to its friendly competitive nature. This makes a fun learning enviroment for people of any age.
# The Design
Last year we designed with the idea of wanting to make something work that looked good. This year, we have been designing with the idea of scaleability. Our database is done in Golang for high scalability, along with MySql and caching stock prices to not overwelm the API resources we are using. I did the initial set up for the database and website while Logan implemented the stocks. Nick then polished database while Logan tied together the webserver with basic front end. Gabe worked on implementing frontend that the team had been sketching while on breaks. *CONTINUE*
# Challenges
Challenges for this Hackathon came from the massive amount of data we needed to handle
* Stock API
  * Aplha Vantage was delayed until the end of the day
  * Morningstar didn't actually have any free API
  * Polygon data was delyed by an hour
  * Finage - what worked in the end - had documentation that differed from how the actual calls worked
* GO Integration
  * We had troubles in getting data pushed from out backend to the frontend through templates
  * MySql being phased out by MariaDB on our Arch and Debian based machines
* Collaboration
  * This year we all had great ideas - and determining which ones to use was important
  * We all have different aesthetic takes for front end work
  * We use three different linux distros and one windows machine - why dockerizing is also on the todo list
# Proud Accomplishments
This is our first hackathon app in which the app completely works. We genuinely plan to use this app with our friend in the calculus class - and are aiming to have actual users for the app. Everything inside of this app is intended to have high scalability, so we have no concerns if the app was to actually take off. We all learned completely new technologies for this app, and contributed to the overall success of the team.
# What We Learned
* Nate
  * While having some minor experience in GO, he never had used it for webservers before. Over the course of the project he worked on the fiber backend for the app.
  * He started to work on the front end - compared to his normal passion for command line applications.
* Gabe
  * Gabe had never coded in GO, but he has gone through commenting the code written by team mates to both further his own understanding and make the project more readable.
  * Gabe built upon his front end knowledge from last HackKU designing our entire front end - and we are proud of how great it looks
* Nick
  * GO Templates
  * MySql with GORM
* Logan
  * GO Templates as well
  * Interacting with a database
