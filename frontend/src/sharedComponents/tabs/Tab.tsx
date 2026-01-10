import { TabButton } from './TabButton';
import { TabPanel } from './TabPanel';
import type { JSX } from 'react';

interface ITabProps {
  activeTabIndex: number;
  currentTabIndex: number;
  buttonProps: {
    label: string;
    onClick: (selectedTabIndex: number) => void;
  };
  panelProps: { children: Maybe<JSX.Element> };
}

export function Tab({
  currentTabIndex,
  activeTabIndex,
  buttonProps,
  panelProps,
}: ITabProps) {
  return (
    <>
      <TabButton
        index={currentTabIndex}
        isActive={currentTabIndex === activeTabIndex}
        onClick={buttonProps.onClick}
        label={buttonProps.label}
      />
      <TabPanel
        activeTabIndex={activeTabIndex}
        children={panelProps.children}
      />
    </>
  );
}
