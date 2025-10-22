export interface modelMarketEntity {
  id: string;
  name: string;
  // provider: ProviderDTO[];
  provider: string;
  address: string;
  input_price: number;
  output_price: number;
  cache_price: number;
  latency: number;
  health_score: number;
  // tokens_per_sec: number;
  last_updated: number;
  subscribed: boolean;
}
