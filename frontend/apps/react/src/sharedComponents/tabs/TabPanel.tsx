import type { JSX } from 'react';
import { getTabId, getTabPanelId } from './tabUtilities';

interface ITabPanelProps {
  children: Maybe<JSX.Element>;
  activeTabIndex: number;
}

export function TabPanel({ children, activeTabIndex }: ITabPanelProps) {
  return (
    <div
      className="grow"
      id={getTabPanelId(activeTabIndex)}
      role="tabpanel"
      tabIndex={0}
      aria-labelledby={getTabId(activeTabIndex)}
    >
      {children}
    </div>
  );
}
