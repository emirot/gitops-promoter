import type { FC } from 'react';
import Card from '@lib/components/Card';
import { type PromotionStrategy } from '@shared/utils/PSData';

interface PromotionStrategyDetailsViewProps {
  strategy: PromotionStrategy;
}

export const PromotionStrategyDetailsView: FC<PromotionStrategyDetailsViewProps> = ({
  strategy,
}) => {
  if (!strategy) return <div>No strategy found</div>;


  //Pass raw data
  const environments = strategy.status?.environments || [];
  return <Card environments={environments} />;
};

export default PromotionStrategyDetailsView; 