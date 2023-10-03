#!/bin/bash
echo "Scraping..."
python AldiScraper.py
python CostcoScraper.py
python PublixScraper.py
python TargetScraper.py