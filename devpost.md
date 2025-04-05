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
