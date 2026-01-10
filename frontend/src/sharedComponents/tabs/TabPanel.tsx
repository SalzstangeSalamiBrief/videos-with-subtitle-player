import { getTabId, getTabPanelId } from './tabUtilities';
import type { JSX } from 'react';

interface ITabPanelProps {
  children: Maybe<JSX.Element>;
  activeTabIndex: number;
}

export function TabPanel({ children, activeTabIndex }: ITabPanelProps) {
  return (
    <div
      className="tab-content bg-base-100 p-4"
      id={getTabPanelId(activeTabIndex)}
      role="tabpanel"
      tabIndex={0}
      aria-labelledby={getTabId(activeTabIndex)}
    >
      {children}
    </div>
  );
}
