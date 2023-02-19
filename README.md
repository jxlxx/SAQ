

# SAQ Scraper

SAQ aka "The Société des alcools du Québec" is a Crown corporation which is responsible for booze manufacturing,
sales, and consumption within the province. And there's no nice, easy-to-use public API if you can believe that :/

This is a simple go program that will scrape the SAQ website after making a query of your choice and save all that
data to a csv file.

To find all results, you can make your query "a" (A is for Alcohol). 
At this moment in time, the "a" query returns 37801 results, which is a bit excessive. 

"Chartreuse" returns 4 results, which may be a better first run. 

# How to use

Build with `go build`.

Run with `./saq_scraper`.

Set `query` to whatever string you wish. Default is "Chartreuse".

Set `language` to either `saq.English` or `saq.Français`. Default is English. 


# Next steps

1. Filtering the results (this means creating a new URL with the proper structure)

2. Scraping individual product pages (I need to know how dry this wine is)

3. Saving to a DB instead of a .csv

4. Updating the aformentioned DB periodically

5. ???

6. Buy the perfect wine for my palate




