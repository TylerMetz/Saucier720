#!/bin/bash
echo "scraping..."
python AldiScraper.py
python CostcoScraper.py
python PublixScraper.py
python TargetScraper.py