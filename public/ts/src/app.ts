const getBtn = document.getElementById("get-trigger") as HTMLButtonElement;
const postBtn = document.getElementById("post-trigger") as HTMLButtonElement;
const putBtn = document.getElementById("put-trigger") as HTMLButtonElement;
const delBtn = document.getElementById("del-trigger") as HTMLButtonElement;
const reqOutput = document.getElementById("req-output") as HTMLTextAreaElement;
const id = document.getElementById("id") as HTMLInputElement;
const statusCode = document.getElementById(
  "status-code"
) as HTMLParagraphElement;

/** Shows the response of the request in the output area and the status of that response
 * @param res The response of the request
 * @function
 * @async
 */
const showRes = async (res: Response) => {
  res.ok
    ? statusCode.classList.add("success")
    : statusCode.classList.remove("success");
  statusCode.innerText = `${res.status} - ${res.statusText}`;
  reqOutput.value = JSON.stringify(await res.json(), null, 2);
};

getBtn.addEventListener("click", async () => {
  const res = await fetch(`/tasks/${id.value}`);
  await showRes(res);
});

postBtn.addEventListener("click", async () => {
  const res = await fetch("/tasks", {
    method: "POST",
    body: reqOutput.value.replace(/(\r\n|\n|\r)/gm, ""),
    headers: new Headers({
      "Content-Type": "application/json",
    }),
  });
  await showRes(res);
});

putBtn.addEventListener("click", async () => {
  const res = await fetch(`/tasks/${id.value}`, {
    method: "PUT",
    body: reqOutput.value.replace(/(\r\n|\n|\r)/gm, ""),
    headers: new Headers({
      "Content-Type": "application/json",
    }),
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
