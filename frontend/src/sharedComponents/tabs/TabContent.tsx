import { getTabId, getTabPanelId } from './tabUtilities';

interface ITabContentProps {
  children: Maybe<JSX.Element>;
  activeTabIndex: number;
}

export function TabContent({ children, activeTabIndex }: ITabContentProps) {
  return (
    <div
      className="flex-grow"
      id={getTabPanelId(activeTabIndex)}
      role="tabpanel"
      tabIndex={0}
      aria-labelledby={getTabId(activeTabIndex)}
    >
      {children}
    </div>
  );
}
