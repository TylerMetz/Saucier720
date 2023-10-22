import logging
import azure.functions as func
import PublixScraper

app = func.FunctionApp()

@app.schedule(schedule="0 0 5 * * 4", arg_name="myTimer", run_on_startup=True,
              use_monitor=False) 
def scrape_publix(myTimer: func.TimerRequest) -> None:
    if myTimer.past_due:
        logging.info('The timer is past due!')

    logging.info('Python timer trigger function executed.')

    deals = PublixScraper.scrape_publix()

    logging.info('Scraping complete.')
