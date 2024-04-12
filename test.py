import requests
import pandas as pd

def get_rsi(symbol, interval='1h', period='14'):
    url = f"https://api.binance.com/api/v3/klines?symbol={symbol}&interval={interval}&limit={period}"
    response = requests.get(url)
    data = response.json()

    df = pd.DataFrame(data, columns=['timestamp', 'open', 'high', 'low', 'close', 'volume', 'close_time', 'quote_asset_volume', 'number_of_trades', 'taker_buy_base_asset_volume', 'taker_buy_quote_asset_volume', 'ignore'])
    df['close'] = pd.to_numeric(df['close'])

    delta = df['close'].diff()
    gain = (delta.where(delta > 0, 0)).rolling(window=int(period)).mean()
    loss = (-delta.where(delta < 0, 0)).rolling(window=int(period)).mean()

    RS = gain / loss
    RSI = 100 - (100 / (1 + RS))

    return RSI.iloc[-1]

# Example usage:
symbol = 'BTCUSDT'  # Bitcoin/USDT 페어
rsi = get_rsi(symbol)
print(f"The RSI for {symbol} is {rsi}")
