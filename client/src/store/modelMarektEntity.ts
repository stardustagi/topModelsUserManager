export interface modelMarketEntity {
  // id: number;
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
  map_node_id: number;
  map_model_id: number;
}
