import { useId, type JSX } from 'react';
import { TabButtons } from './TabButtons';
import { TabPanel } from './TabPanel';

export interface ITab {
  label: string;
  content: Maybe<JSX.Element>;
}

export interface ITabsProps {
  tabs: ITab[] | undefined;
  label: string;
  activeTabIndex: number | number;
  onChangeTab: (index: number) => void;
}

export function Tabs({
  activeTabIndex,
  label,
  onChangeTab,
  tabs = [],
}: ITabsProps) {
  const labelId = useId();

  if (!tabs.length) {
    console.warn('No tabs provided');
    return null;
  }

  const activeTab = tabs[activeTabIndex];

  return (
    <section aria-labelledby={labelId} className="h-full pb-8">
      <h1 id={labelId} className="sr-only">
        {label}
      </h1>
      <div role="tablist" className="flex h-full flex-col gap-8">
        <TabButtons
          tabs={tabs}
          activeTabIndex={activeTabIndex}
          onChangeTab={onChangeTab}
        />
        <TabPanel activeTabIndex={activeTabIndex}>{activeTab.content}</TabPanel>
      </div>
    </section>
  );
}
