export function Navigation() {
  return <h1>Navigation</h1>;
  // const { folderId } = useParams({ strict: false });
  // const { fileTrees } = useContext(FileTreeContext);
  // const activeFolders = getFoldersInActiveTree(fileTrees, folderId);

  // return (
  //   <nav className={styles.navigation}>
  //     <Menu<IFileTreeDto>
  //       itemKey="id"
  //       items={fileTrees}
  //       activeItemIds={activeFolders.map((a) => a.id)}
  //       onRenderMenuItem={(item: IFileTreeDto) => {
  //         const isActive = activeFolders.some((a) => a.id === item.id);
  //         return (
  //           <NavigationItem
  //             item={item}
  //             isActive={isActive}
  //             hasChildren={Boolean(item.children?.length)}
  //           />
  //         );
  //       }}
  //     />
  //   </nav>
  // );
}
