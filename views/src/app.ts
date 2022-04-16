const getBtn = document.getElementById("get-trigger") as HTMLButtonElement;
const postBtn = document.getElementById("post-trigger") as HTMLButtonElement;
const putBtn = document.getElementById("put-trigger") as HTMLButtonElement;
const delBtn = document.getElementById("del-trigger") as HTMLButtonElement;
const reqOutput = document.getElementById("req-output") as HTMLTextAreaElement;
const id = document.getElementById("id") as HTMLInputElement;
const statusCode = document.getElementById(
  "status-code"
) as HTMLParagraphElement;

const showRes = async (res: Response) => {
  res.ok
    ? statusCode.classList.add("success")
    : statusCode.classList.remove("success");
  statusCode.innerText = `${res.status} - ${res.statusText}`;
  reqOutput.value = JSON.stringify(await res.json(), null, 2);
};

const REPLACE_REGEX = /(\r\n|\n|\r)/gm;

getBtn.addEventListener("click", async () => {
  const res = await fetch(`/tasks/${id.value}`);
  await showRes(res);
});

postBtn.addEventListener("click", async () => {
  const res = await fetch("/tasks", {
    method: "POST",
    body: reqOutput.value.replace(REPLACE_REGEX, ""),
    headers: {
      "Content-Type": "application/json",
    },
  });
  await showRes(res);
});

putBtn.addEventListener("click", async () => {
  const res = await fetch(`/tasks/${id.value}`, {
    method: "PUT",
    body: reqOutput.value.replace(REPLACE_REGEX, ""),
    headers: {
      "Content-Type": "application/json",
    },
  });
  await showRes(res);
});

delBtn.addEventListener("click", async () => {
  const deleteConfirmed = id.value
    ? confirm("Are you sure you want to delete this task?")
    : confirm("Are you sure you want to delete all tasks?");
  if (deleteConfirmed) {
    const res = await fetch(`/tasks/${id.value}`, {
      method: "DELETE",
    });
    await showRes(res);
  }
});
