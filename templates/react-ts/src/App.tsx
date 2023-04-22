import { useState } from "react";

type ColoredSpanType = {
  color: string;
  children: any;
};
function ColoredSpan({ color, children }: ColoredSpanType) {
  return <span className={`color_${color}`}>{children}</span>;
}

function App() {
  const [count, setCount] = useState<number>(0);
  const template = "${TEMPLATE}";
  const app_name = "${APP_NAME}";
  return (
    <div className="App">
      <pre>
        <code>
          <ColoredSpan color={"red"}>webdev</ColoredSpan>
          <ColoredSpan color={"default"}>gen</ColoredSpan>
          <ColoredSpan color={"green"}>{template}</ColoredSpan>
          <ColoredSpan color={"green"}>{app_name}</ColoredSpan>
        </code>
      </pre>
      <button type="button" onClick={() => setCount((count) => count + 1)}>
        count is: {count}
      </button>
    </div>
  );
}

export default App;
