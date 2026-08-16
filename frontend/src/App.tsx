import { useState, type PropsWithChildren } from "react";
import "./App.css";
import Header from "./components/header";
import type { PageType } from "./types";
import useSearchParams from "./hooks/useSearchParams";
import Mixer from "./components/pages/mixer";
import Scenes from "./components/pages/scenes";
import Performance from "./components/pages/performance";
import Samples from "./components/pages/samples";
import Page from "./components/page";
import { SocketProvider } from "./providers/SocketProvider";
import { X32SProvider } from "./providers/X32Provider";

export default function App() {
  const pages: PageType[] = [
    {
      id: "mixer",
      name: "Mixer",
      page: <Mixer />,
    },
    {
      id: "scenes",
      name: "Scenes",
      page: <Scenes />,
    },
    {
      id: "performance",
      name: "Performance",
      page: <Performance />,
    },
    {
      id: "samples",
      name: "Samples",
      page: <Samples />,
    },
  ];

  const { searchParams, setSearchParams } = useSearchParams();

  const [page, setPage] = useState(() => {
    const pageQuery = searchParams.get("page");
    if (pageQuery != null) {
      const currentPage = pages.find((p) => p.id == pageQuery);
      if (currentPage) {
        setSearchParams("page", currentPage.id);
        return currentPage;
      }
    }

    const currentPage = pages[0];
    setSearchParams("page", currentPage.id);
    return currentPage;
  });

  const navigatePage = (pageId: string) => {
    const page = pages.find((p) => p.id == pageId);
    if (page) {
      setSearchParams("page", page.id);
      setPage(page);
    }
  };

  return (
    <AppProviders>
      <Header pages={pages} currentPage={page} setPage={navigatePage} />
      {pages.map((p) => (
        <Page hidden={p !== page} key={p.id}>
          {p.page}
        </Page>
      ))}
    </AppProviders>
  );
}

function AppProviders({ children }: PropsWithChildren) {
  return (
    <SocketProvider>
      <X32SProvider>{children}</X32SProvider>
    </SocketProvider>
  );
}
