import { TabButton } from './TabButton';
import { ITab } from './Tabs';

interface ITabButtonsProps {
  tabs: ITab[];
  activeTabIndex: number;
  onChangeTab: (index: number) => void;
}

export function TabButtons({
  tabs,
  activeTabIndex,
  onChangeTab,
}: ITabButtonsProps) {
  return (
    <div role="presentation" className="flex gap-1">
      {tabs.map((tab, index) => {
        const isActive = index === activeTabIndex;
        return (
          <TabButton
            key={index}
            isActive={isActive}
            onClick={() => onChangeTab(index)}
            label={tab.label}
            index={index}
          />
        );
      })}
    </div>
  );
}
