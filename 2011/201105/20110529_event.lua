EVENT_SAMPLE = 1000
RegisterEvent("EventHandler")
function EventHandler(id, ...)
  if id == EVENT_SAMPLE then
    print("Sample Event!")
  end
end