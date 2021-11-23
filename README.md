# goinvest
investing.com API for stock screener written in Golang

текущий затык:
1. investing.com не имеют своего api, а значит надо использовать crawl'ер;
2. банят после нескольких запросов. не уверен, что по IP, возможно по user agent;
3. рендерят страницы через JS (омайгад), поэтому надо использовать crawl'ер на основе chrome (омайгад);
