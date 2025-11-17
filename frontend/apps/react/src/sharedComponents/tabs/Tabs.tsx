import type { Maybe } from '@videos-with-subtitle-player/core';
import { useId, type JSX } from 'react';
import { Tab } from './Tab';

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
      <div role="tablist" className="tabs tabs-lift">
        {tabs.map((tab, index) => (
          <Tab
            key={index}
            activeTabIndex={activeTabIndex}
            currentTabIndex={index}
            buttonProps={{ label: tab.label, onClick: onChangeTab }}
            panelProps={{ children: activeTab.content }}
          />
        ))}
        {/* <div role="tablist" className="tabs tabs-box"> */}
        {/* <TabButtons
          tabs={tabs}
          activeTabIndex={activeTabIndex}
          onChangeTab={onChangeTab}
        />
        <TabPanel activeTabIndex={activeTabIndex}>{activeTab.content}</TabPanel> */}
      </div>
    </section>
  );
}
