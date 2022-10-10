import Head from "next/head";
import { PropsWithChildren } from "react";

export default function Layout({ children }: PropsWithChildren<{}>) {
  return (
    <div className={"container"}>
      <Head>
        <title>IoT App</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className={"main"}>{children}</main>
      <footer></footer>
    </div>
  );
}
