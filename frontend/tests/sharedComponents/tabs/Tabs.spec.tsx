import type { ITab } from '$sharedComponents/tabs/Tabs';
import { Tabs } from '$sharedComponents/tabs/Tabs';
import { render, screen, within } from '@testing-library/react';
import { describe, expect, test } from 'vitest';

describe('Tabs', () => {
  test('Renders nothing if no tabs are provided', () => {
    const { container } = render(
      <Tabs label="Test" activeTabIndex={0} onChangeTab={() => {}} tabs={[]} />,
    );
    expect(container.innerHTML).toBe('');
  });

  test('Render one tab with content', () => {
    const sectionLabel = 'Sectionlabel';
    const tabContent = 'tabContent';
    const singleTab: ITab = { label: 'Test', content: <div>{tabContent}</div> };
    render(
      <Tabs
        label={sectionLabel}
        activeTabIndex={0}
        onChangeTab={() => {}}
        tabs={[singleTab]}
      />,
    );

    expect(screen.getByText(sectionLabel)).toBeInTheDocument();
    const tablist = screen.getByRole('tablist');
    const tabs = within(tablist).getAllByRole('tab');
    expect(tabs.length).toBe(1);
    expect(tabs[0]).toHaveTextContent(singleTab.label);
    const tabpanel = within(tablist).getByRole('tabpanel');
    expect(tabpanel).toHaveTextContent(tabContent);
  });

  describe('Render multiple tabs', () => {
    const differentTabsTestCases = Array.from({ length: 5 }).map<{
      label: string;
      content: string;
    }>((_, i) => ({ label: `Tab ${i}`, content: `Content ${i}` }));

    test.each(differentTabsTestCases)(
      'Render multiple tabs with active tab $label',
      ({ label, content }) => {
        const i = differentTabsTestCases.findIndex(
          (tab) => tab.label === label && tab.content === content,
        );
        render(
          <Tabs
            label="SectionLabel"
            activeTabIndex={i}
            onChangeTab={() => {}}
            tabs={differentTabsTestCases.map(
              ({ label: testLabel, content: testContent }) => ({
                label: testLabel,
                content: <div>{testContent}</div>,
              }),
            )}
          />,
        );

        const tablist = screen.getByRole('tablist');
        const tabs = within(tablist).getAllByRole('tab');
        expect(tabs.length).toBe(differentTabsTestCases.length);
        tabs.forEach((tab, j) => {
          expect(tab).toHaveTextContent(differentTabsTestCases[j].label);
          expect(tab).toHaveAttribute(
            'aria-selected',
            j === i ? 'true' : 'false',
          );
        });

        const tabpanel = within(tablist).getByRole('tabpanel');
        expect(tabpanel).toHaveTextContent(content);
      },
    );
  });
});
