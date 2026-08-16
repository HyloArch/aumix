import "./Header.css";
import type { PageType } from "../../types";
import { clsx } from "../../util";
import StatusIcon from "../status-icon";

type HeaderProps = {
  pages: PageType[];
  currentPage: PageType;
  setPage: (page: string) => void;
};

export default function Header({ pages, currentPage, setPage }: HeaderProps) {
  return (
    <header>
      <div className="title">AU Mix</div>
      <div className="tabs">
        {pages.map((page) => (
          <div
            key={page.id}
            className={clsx("tab", { selected: page === currentPage })}
            onClick={() => setPage(page.id)}
          >
            {page.name}
          </div>
        ))}
      </div>
      <StatusIcon />
    </header>
  );
}
